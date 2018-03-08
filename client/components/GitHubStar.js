import React from "react";

class GitHubStar extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            starred: true,
            nStars: 42
        };
    }

    componentDidMount() {
        this.timerID = setInterval(() => this.tick(), 2000);
    }

    tick() {
        this.setState({
            nStars: Math.floor(Math.random() * 2400) + 1
        });
    }

    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    render() {
        return (
            <div style={gitHubStyles.button}>
                <div style={gitHubStyles.leftSide}>
                    <div style={gitHubStyles.star}>
                        <svg
                            aria-hidden="true"
                            height="16"
                            version="1.1"
                            viewBox="0 0 14 16"
                            width="14"
                        >
                            <path
                                fillRule="evenodd"
                                d="M14 6l-4.9-.64L7 1 4.9 5.36 0 6l3.6 3.26L2.67 14 7 11.67 11.33 14l-.93-4.74z"
                            />
                        </svg>
                    </div>
                    <span style={gitHubStyles.text}>
                        {this.state.starred ? "Unstar" : "Star"}
                    </span>
                </div>
                <div style={gitHubStyles.rightSide}>
                    <span style={gitHubStyles.text}>{this.state.nStars}</span>
                </div>
            </div>
        );
    }
}

// Inline styles because we don't need these styles anywhere else
const gitHubStyles = {
    star: {
        marginRight: "4px"
    },
    button: {
        display: "inline-flex",
        borderRadius: "4px",
        fontSize: "12px",
        color: "#24292e",
        fontWeight: 600,
        lineHeight: 1,
        whiteSpace: "nowrap",
        verticalAlign: "middle",
        backgroundRepeat: "repeat-x",
        backgroundPosition: "-1px -1px",
        backgroundSize: "110% 110%",
        border: "1px solid rgba(27,31,35,0.2)",
        borderRadius: "0.25em"
    },
    leftSide: {
        display: "flex",
        alignItems: "center",
        color: "#24292e",
        backgroundColor: "#eff3f6",
        backgroundImage: "linear-gradient(-180deg, #fafbfc 0%, #eff3f6 90%)",
        padding: "5px 10px",
        borderRight: "1px solid rgba(27,31,35,0.2)"
    },
    rightSide: {
        display: "flex",
        alignItems: "center",
        padding: "5px 10px"
    },
    text: {
        fontFamily: "-apple-system, sans-serif"
    }
};

export default GitHubStar;
