import React from "react";

const Caption = ({ figureNumber, figureDescription }) => (
    <div className="caption">
        <span className="italic">{`Figure ${figureNumber}. `}</span>
        <span>{figureDescription}</span>
    </div>
);

export default Caption;
