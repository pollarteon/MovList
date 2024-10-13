const express = require("express");
const app = express();
const port = 5000;

app.post("/movies",()=>{
    console.log("Movie added to Database!");
})


app.listen(port,()=>{
    console.log("Listening on port 5000");
})
