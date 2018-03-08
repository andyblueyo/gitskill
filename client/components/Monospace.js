import React from "react";

class Monospace extends React.Component {
    render() {
        return (
            <p style={{ fontFamily: "monospace", color: "#4e5052" }}>
                {this.props.children}
            </p>
        );
    }
}

export default Monospace;
