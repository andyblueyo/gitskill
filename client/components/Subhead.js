import React from "react";

const Subhead = ({ authors }) => {
    let formattedAuthors = [applyStyles(true, authors[0])];

    for (let i = 1; i < authors.length; i++) {
        formattedAuthors.push(applyStyles(false, authors[i]));
    }

    return (
        <div className="measure mx-auto">
            <span className="subhead">By {formattedAuthors}</span>
        </div>
    );
};

const applyStyles = (isFirst, author) => {
    return (
        <span>
            {isFirst ? "" : ", "}
            <a
                className="black"
                href={author.url}
                target="_blank"
                key={author.name}
            >
                {author.name}
            </a>
        </span>
    );
};

module.exports = Subhead;
