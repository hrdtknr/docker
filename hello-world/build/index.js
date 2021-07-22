"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const body_parser_1 = require("body-parser");
const app = express_1.default();
app.use(body_parser_1.urlencoded({ extended: true }));
app.get('/', (req, res) => {
    res.send('Hello World');
});
app.post('/', (req, res) => {
    console.log(req.body);
    res.send({ 'post': req.body });
});
app.listen(3000);
//# sourceMappingURL=index.js.map