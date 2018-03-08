import React from "react";

class List extends React.Component {
    render() {
        console.log(this.props.items);
        return (
            <ul>
                {this.props.items.map((item, i) => <li key={i}>{item}</li>)}
            </ul>
        );
    }
}

export default List;
