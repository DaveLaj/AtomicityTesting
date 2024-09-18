# SQL Isolation Levels

## Applying Isolation Levels To SQL Transactions in Go

### [https://pkg.go.dev/database/sql#Conn.BeginTx]

```Go
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
```

### Example

```Go
sqlOpts := &sql.TxOptions{
    Isolation: sql.LevelReadCommitted // Read sql.Constants for more isolation levels
}

tx, err := db.BeginTx(ctx, sqlOpts)
if err != nil {
    // handle error
}
```

## Explanation

The four isolation levels defined in the SQL standard, in increasing order of isolation attained for a given transaction, are `READ UNCOMMITTED`, `READ COMMITTED`, `REPEATABLE READ`, and `SERIALIZABLE`.

## `READ UNCOMMITTED`

`READ UNCOMMITTED` is level 0 locking, which means that no shared locks are issued and no exclusive locks are honored. This isolation level allows dirty reads, i.e., where a transaction may see uncommitted changes made by some other transaction. This means values in the data can be changed and rows can appear or disappear in the data set before the transaction completes.

**The tl;dr: In general, don’t use `READ UNCOMMITTED`.** The one and only time it ever makes sense is when you are reading data that will never be modified in any way. For example, if you need to scan a large volume of data to generate high-level analytics or summaries, and absolute moment-of-query accuracy is not critical.

## `READ COMMITTED`

`READ COMMITTED` is the weakest isolation level implemented by most databases. It is the default isolation level in most Postgres databases, and older SQL databases in general.

This level of isolation allows the transaction to see different versions of the data in different statements. `READ COMMITTED` restricts the reader from seeing any intermediate, uncommitted, ‘dirty’ read during a transaction and guarantees that any data read was committed at the moment it was read. This is what makes `READ COMMITTED` a popular default option, because this effectively eliminates most transaction retry errors (when a transaction cannot be completed and so is aborted and tried again). In most cases where a transaction with a higher level of isolation would return a retry error, a `READ COMMITTED` transaction can retry only the current statement — and this retry can be handled invisibly on the server.

`READ COMMITTED` is a good choice if:

- You are migrating an application that was originally developed for a different database that used this isolation level.

- It is difficult to modify the application to handle transaction retry errors

- Consistently low latency is important for the application and full transaction retries are too expensive.

One use case example where `READ COMMITTED` is the best option would be a retail application that displays the available inventory of products. It’s a good balance between consistency and concurrency in situations where it’s important to see only committed data (so, no dirty reads) but minor inconsistencies between reads (like non-`REPEATABLE READ`s) are acceptable. When displaying available inventory to customers, most ecommerce applications can afford momentary slight inaccuracies caused by concurrent updates, as long as the data is not completely out of date or showing uncommitted data.

It is important to note that `READ COMMITTED` makes no promise whatsoever that, if the transaction re-issues the read, it will find the same data. In this level of isolation, data is free to change in the interval after it was read but before the transaction completes — leading to the database anomaly known as non`REPEATABLE READ`s

## `REPEATABLE READ`

`REPEATABLE READ` is an intermediate level of isolation that inherits some of the drawbacks of both its weaker and stronger neighbors. As on the weaker side of the isolation spectrum, explicit locks (including SELECT FOR UPDATE) are still sometimes required to prevent data anomalies. And, as in the stronger isolation levels, application-visible transaction retry errors can occur. So why would you ever use it?

Because `REPEATABLE READ` is the ideal isolation level for read-only transactions.

Using the lower isolation level of `READ COMMITTED` for a series of read-only statements is inefficient, because it will essentially run each of those statements in their own individual, single-statement transaction. `SERIALIZABLE` isolation is also inefficient because it will do extra work to detect read/write conflicts that cannot possibly exist in a read-only transaction.

So, for example, `REPEATABLE READ` would be a good choice for a financial app that calculates the total balance of a user’s accounts. This isolation level ensures that if a row is read twice in the same transaction, it will return the same value each time, preventing the non`REPEATABLE READ` anomaly mentioned above. In this case, when calculating a user’s total balance, `REPEATABLE READ` guarantees that the balance does not change during the calculation process due to concurrent updates.

Again: **`REPEATABLE READ` is your friend in read-only transactions**. For read-write transactions, though, choose either `SERIALIZABLE` or `READ COMMITTED` isolation levels.

## `SERIALIZABLE`

`SERIALIZABLE` is the highest level of isolation. Transactions are completely isolated from each other, effectively serializing access to the database to prevent dirty reads, non-`REPEATABLE READ`s, and phantom reads.

The good news is that it prevents any and all isolation anomalies: `SERIALIZABLE` isolation provides a rigorous guarantee that each transaction sees a wholly consistent view of the database, even when there are concurrent transactions. The not-so-good news is that, since it is the most conservative in detecting potential interference between transactions, this isolation produces more transaction retry errors…And the cost of redoing complex transactions can be significant.

Still, you go `SERIALIZABLE` when your application’s transaction logic is sufficiently complex that using `READ COMMITTED` could result in anomalies. `SERIALIZABLE` isolation is absolutely necessary when a transaction executes several successive commands that all must see identical views of the database.

`SERIALIZABLE` is a good choice if:

Data accuracy is paramount and you don’t want to risk anomalies due to missing FOR UPDATE locks.

You are able to use abstractions in your application that let you avoid repeating the retry loop code throughout the app.

You have a retry loop at another level that obviates the need for retries in the app server. For example, a mobile app will often retry failed API calls to cover for flaky networks, and this retry can also cover issues related to serializability.

A typical use case for when `SERIALIZABLE` is the best option would be a banking or financial payments system that processes transactions that transfer money between accounts.

As the highest level of isolation, it’s best suited for scenarios requiring strict consistency and where the integrity of each transaction is critical. After all, when transferring money between accounts, it’s essential that the amount that leaves Account A arrives intact at Account B with no other transactions interfering. Data anomalies are never good, but they are 10x worse when causing problems like incorrect or missing funds in your bank account
