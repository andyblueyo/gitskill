#!/usr/bin/env node
const users = require("./data/users.json");
const repos = require("./data/repos.json");

setTimeout(function() {
    let total = 0;
    let totalDiff = 0;
    users.forEach(u => {
        if (u.userType.toLowerCase() === "user") {
            let name = u.githubUsername;
            let count = u.publicRepos;
            let orgRepos = repos.filter((repo) => {
                return repo.ownerName === name;
            });
            if (count !== orgRepos.length && count - 1 !== orgRepos.length) {
                console.log(`issue with ${name}, expected ${count} but got ${orgRepos.length}`);
                total += 1;
                totalDiff += (count - orgRepos.length);
            }
        }
    });
    console.log('total users not equal:', total);
    console.log('total repo diff:', totalDiff);
    process.exit();
}, 1000);

