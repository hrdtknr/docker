import express from 'express';
import { urlencoded } from 'body-parser';

const app: express.Express = express();

app.use(urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.send('Hello World');
})

app.post('/', (req, res) => {
  console.log(req.body)
  res.send({'post':req.body});
})

app.listen(3000);