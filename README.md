# golang-javascript-async-comparison
A set of async patterns implemented in both golang and javascript. Serves as a translation between the two languages.

## To Run
- Javascript: `node example.mjs`

## Javascript Notes
- '.mjs' extention is used so top level await can be used.
- Fetch api is stil experimental in nodejs 18.10.


### TODO: examples
- Queue which only allows x number of http calls to be made at the same time (http-agent).
- In memory batch which fires off after certain number of things have been queued.
- In memory batch which fires off after time limit has exceeded.
- Producer/consumer.
- Fire and forget.