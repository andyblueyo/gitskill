#!/usr/bin/env node
const users = require("./data/users.json");
const repos = require("./data/repos.json");

setTimeout(function() {
    let total = 0;
    users.forEach(u => {
        let name = u.githubUsername;
        let count = u.publicRepos;
        let orgRepos = repos.filter((repo) => {
            return repo.ownerName === name;
        });
        if (count !== orgRepos.length) {
            console.log(`issue with ${name}, expected ${count} but got ${orgRepos.length}`)
            total += 1;
        }
    });
    console.log(total);
    process.exit();
}, 5000);


