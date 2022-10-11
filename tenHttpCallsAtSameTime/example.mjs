async function makeHttpCall() {
    try {
        const response = await fetch('https://jsonplaceholder.typicode.com/todos/1')
        const json = await response.json()
        console.log(json)
    }
    catch (error) {
        console.error(error)
    }

}

const promises = []
for(let i = 0; i < 10; i++) {
    promises.push(makeHttpCall())
}

await Promise.allSettled(promises)