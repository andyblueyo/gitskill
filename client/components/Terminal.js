import moment from "moment";
import React from "react";

const TerminalHeader = props => (
    <div style={styles.terminalHeader}>
        <div style={styles.terminalHeaderInner}>
            <div style={styles.buttonContainer}>
                <div style={styles.button} />
                <div style={styles.button} />
                <div style={styles.button} />
            </div>
            {props.children}
            <div style={styles.fauxButtonContainer} />
        </div>
    </div>
);

const TerminalBody = props => (
    <div style={styles.terminalBody}>
        <div style={styles.terminalBodyInner}>{props.children}</div>
    </div>
);

class Terminal extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            output: [generateLine()]
        };
    }

    componentDidMount() {
        this.timerID = setInterval(() => this.tick(), 1000);
    }

    tick() {
        if (this.state.output.length !== 12) {
            this.setState({
                output: [...this.state.output, generateLine()]
            });
        } else {
            this.setState({
                output: [generateLine()]
            });
        }
    }

    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    render() {
        const formattedOutput = this.state.output.map(line => (
            <div>{line}</div>
        ));

        return (
            <div id="terminal">
                <TerminalHeader>
                    <span>evan ― -bash ―</span>
                </TerminalHeader>
                <TerminalBody>{formattedOutput}</TerminalBody>
            </div>
        );
    }
}

function generateLine() {
    return `[${getTime()}] Fetching repository ${makeId()}`;
}

// Source: https://stackoverflow.com/questions/1349404/generate-random-string-characters-in-javascript
function makeId() {
    return Math.random()
        .toString(36)
        .substring(2, 8);
}

function getTime() {
    return moment().format("HH:mm:ss");
}

const styles = {
    terminalHeader: {
        background:
            "linear-gradient(0deg, rgb(216, 216, 216), rgb(236, 236, 236))",
        borderTop: "1px solid white",
        color: "rgba(0, 0, 0, 0.7)",
        fontFamily: "Helvetica, sans-serif",
        fontSize: "13px",
        textAlign: "center",
        lineHeight: "22px",
        height: "22px",
        borderTopLeftRadius: "5px",
        borderTopRightRadius: "5px"
    },
    terminalHeaderInner: {
        display: "flex",
        justifyContent: "space-between"
    },
    buttonContainer: {
        display: "flex",
        alignItems: "center"
    },
    button: {
        height: "12px",
        width: "12px",
        borderRadius: "50%",
        marginLeft: "4px",
        marginRight: "4px",
        border: "1px solid rgba(0, 0, 0, 0.2)",
        backgroundColor: "#8c8c91"
    },
    fauxButtonContainer: {
        width: "52px"
    },
    terminalBody: {
        backgroundColor: "#323232",
        color: "white",
        fontFamily: "Menlo, monospace",
        fontSize: "13px",
        whiteSpace: "pre-wrap",
        borderBottomLeftRadius: "5px",
        borderBottomRightRadius: "5px",
        paddingBottom: "56.25%",
        position: "relative",
        overflow: "hidden"
    },
    terminalBodyInner: {
        position: "absolute",
        top: "4px",
        right: "4px",
        bottom: "4px",
        left: "4px"
    }
};

export default Terminal;
