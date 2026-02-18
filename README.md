# Wallet Microservice

## Locking Mechanism

I choose optimistic locking mechanism for the db select and update. From the sentence "A user has a balance in a digital wallet application," I got the information that I am creating a digital wallet application back-end service. Since digital wallet is most likely accessed by a user instead of a company, I conclude that there won't be many volume for the withdrawals.

Though there's also an argument that the get balance might be outdated, like 50ms outdated if the user just withdrawn and the read take a bit longer than usual, it's still okay. Since the user most likely will notice and pull the refresh button in their front-end, or get the accurate balance when they log in to their account again.
