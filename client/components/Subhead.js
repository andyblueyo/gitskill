import React from "react";

const Subhead = ({ authors }) => {
    let formattedAuthors = [];

    for (let i = 0; i < authors.length; i++) {
        switch (i) {
            case 0:
                formattedAuthors.push(applyStyles(authors[i], true, false));
                break;
            case authors.length - 1:
                formattedAuthors.push(applyStyles(authors[i], false, true));
                break;
            default:
                formattedAuthors.push(applyStyles(authors[i], false, false));
        }
    }

    return (
        <div className="measure mx-auto">
            <span className="subhead">By {formattedAuthors}</span>
        </div>
    );
};

const applyStyles = (author, isFirst, isLast) => {
    let predecessor = "";

    if (isLast) {
        predecessor = " and ";
    } else if (!isFirst) {
        predecessor = ", ";
    }

    return (
        <span>
            {predecessor}
            <a
                className="black bold sans-serif"
                href={author.url}
                target="_blank"
                key={author.name}
            >
                {author.name}
            </a>
        </span>
    );
};

export default Subhead;
