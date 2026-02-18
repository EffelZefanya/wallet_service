# Wallet Microservice

## Decisions in Project Building

## Locking Mechanism

I choose optimistic locking mechanism for the db select and update. From the sentence "A user has a balance in a digital wallet application," I got the information that I am creating a digital wallet application back-end service. Since digital wallet is most likely accessed by a user instead of a company, I conclude that there won't be many volume for the withdrawals.

Though there's also an argument that the get balance might be outdated, like 50ms outdated if the user just withdrawn and the read take a bit longer than usual, it's still okay. Since the user most likely will notice and pull the refresh button in their front-end, or get the accurate balance when they log in to their account again.

## GRPC vs REST API

I chose to use REST API instead of GRPC because from the statement of "A user has a balance in a digital wallet application and wants to perform transactions, including withdrawals and balance inquiries." Since an application is used, I conclude that this back-end will be called right away by the front-end.

I do not use a GRPC because it's usually used in service and the application as of now can be finished with a monolith architecture. Using a GRPC will only lead to unnecessary complexities in the future.

## No Password in User Table

Since this project focuses more on the wallet service, I didn't put "password" as a column in the user table. Making a "password" column will require me to create a login and register function, which is outside of this project's scope. This also lead to quicker development process.

## Not every Table has updated_at

To quicken development speed, updated_at are only set in wallet table. It is the only one that will get updated in this case.
