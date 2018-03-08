import React from "react";

const Desktop = ({ children }) => {
    return <div className="hide-xs hide-sm hide-md">{children}</div>;
};

export default Desktop;
