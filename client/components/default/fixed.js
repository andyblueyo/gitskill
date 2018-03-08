import React from "react";

class Fixed extends React.PureComponent {
    render() {
        return <div className="fixed">{this.props.children}</div>;
    }
}

export default Fixed;
