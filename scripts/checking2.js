#!/usr/bin/env node
const users = require("./data/users.json");
// const repos = require("./data/repos.json");

setTimeout(function () {
    let orgMap = {};
    // orgs
    users.filter(u => {
        return u.userType.toLowerCase() === "organization";
    }).map((org) => {
        return org.githubUsername;
    }).forEach(name => {
        orgMap[name] = 0;
    });

    // user orgs
    let doubleArrayOfOrgs = users.filter(u => {
        return u.userType.toLowerCase() !== "organization";
    }).map((user) => {
        return user.orgs;
    });
    doubleArrayOfOrgs.forEach(oArr => {
        if(oArr) {
            oArr.forEach(org => {
                if (orgMap.hasOwnProperty(org)) {
                    orgMap[org] = orgMap[org] + 1;
                }
            })
        }
    });

    Object.keys(orgMap).forEach((o) => {
        if (orgMap[o] === 0) {
            console.log("not yet found org:", o)
        }
    });
}, 1000);


