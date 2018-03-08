import React from "react";

class Heading extends React.Component {
    render() {
        const { is, children } = this.props;

        switch (is) {
            case "h1":
                return <h1>{children}</h1>;
            case "h2":
                return <h2>{children}</h2>;
            case "h3":
                return <h3>{children}</h3>;
            case "h4":
                return <h4>{children}</h4>;
            case "h5":
                return <h5>{children}</h5>;
            case "h6":
                return <h6>{children}</h6>;
            default:
                return null;
        }
    }
}

export default Heading;
