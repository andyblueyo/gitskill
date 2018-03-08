import React from "react";
import _ from "lodash";

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
                <table>
                    <thead>{thead}</thead>
                    <tbody>{tbody}</tbody>
                </table>
            </div>
        );
    }
}

export default TableComponent;
