import React from "react";

const Desktop = ({ children }) => {
    return <div className="hide-xs hide-sm hide-md hide-lg">{children}</div>;
};

export default Desktop;
