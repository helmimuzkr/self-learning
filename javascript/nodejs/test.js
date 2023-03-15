import { mkdirSync, existsSync, writeFileSync, readFileSync } from "fs";
import readline from "readline";
import { json } from "stream/consumers";

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

const dirPath = "./data";
if (!existsSync(dirPath)) {
    mkdirSync(dirPath);
}

const filePath = "./data/user.json";
if (!existsSync(filePath)) {
    writeFileSync(filePath, "[]", "utf8");
}

rl.question("Ketikan Nama Anda : ", (name) => {
    rl.question("Ketikan Umur Anda : ", (age) => {
        const user = { name, age };

        const post = JSON.parse(readFileSync("./data/user.json", "utf8"));

        post.push(user);

        writeFileSync("./data/user.json", JSON.stringify(post));

        rl.close();
    });
});
