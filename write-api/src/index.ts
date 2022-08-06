import express from 'express';
import router from "./routes";

const app = express();

app.use(express.json())
app.use(express.urlencoded({ extended: true }));

app.use('/', router);

const port = process.env.PORT || 3000;

app.listen(port);
console.log('Express WebApi listening on port ' + port);