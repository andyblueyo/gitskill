import React from "react";
import _ from "lodash";

import Caption from "./Caption";

class TableComponent extends React.Component {
    componentDidMount() {
        console.log(this.props.data);
    }
    render() {
        let thead = (
            <tr>{this.props.columns.map(column => <th>{column}</th>)}</tr>
        );

        let tbody = this.props.data.map((row, i) => (
            <tr>{_.map(row, value => <td>{value}</td>)}</tr>
        ));

        return (
            <div style={{ overflowX: "auto" }}>
                <div>
                    <table>
                        <thead>{thead}</thead>
                        <tbody>{tbody}</tbody>
                    </table>
                </div>
                <Caption
                    figureNumber={this.props.figureNumber}
                    figureDescription={this.props.figureDescription}
                />
            </div>
        );
    }
}

export default TableComponent;
