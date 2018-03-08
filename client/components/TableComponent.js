import React from "react";
import _ from "lodash";

import Caption from "./Caption";

class TableComponent extends React.Component {
    componentDidMount() {
        // console.log(this.props.data);
    }
    render() {
        let thead = (
            <tr>
                {this.props.columns.map((column, i) => (
                    <th key={`column_${i}`}>{column}</th>
                ))}
            </tr>
        );

        let tbody = this.props.data.map((row, i) => {
            return (
                <tr key={`row_${i}`}>
                    {_.map(row, (value, j) => (
                        <td key={`tableKey_${i}_${j}`}>{value}</td>
                    ))}
                </tr>
            );
        });

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
