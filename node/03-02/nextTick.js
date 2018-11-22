function func(callback) {
    // Call callback function which is passed as a parameter using nextTick
    process.nextTick(callback, "callback!!");
}

// try & catch is being run on the same thread
// Therefore, the exception process is not able to be executed.
try {
    func((param) => {
        a.a = 0;
    });
} catch (e) {
    console.log("exception!!");
}
