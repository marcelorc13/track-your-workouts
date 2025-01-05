import { app } from "./app";
import  MongoDatabase  from "./db/mongodb";
import { configDotenv } from "dotenv";
configDotenv()

const port = process.env.PORT

MongoDatabase.connect()
app.listen(port, () => console.log(`O app está rodando em http://localhost:${port}`))
