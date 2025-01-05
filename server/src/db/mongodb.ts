import mongoose from "mongoose"
import { configDotenv } from "dotenv"
configDotenv()

class MongoDatabase {
    private uri = `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0.lt2hq.mongodb.net/Express-Node-API?retryWrites=true&w=majority&appName=Cluster0`

    public connect = () => {
        mongoose.connect(this.uri)
            .then(() => {
                console.log("Conected to the Database")
            }).catch(() => {
                console.log("Conection to DB failed")
            })
    }
}

export default new MongoDatabase