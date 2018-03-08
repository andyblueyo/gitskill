import React from "react";

import Flex from "./Flex";
import Blockquote from "./Blockquote";

class SidePanel extends React.Component {
    render() {
        let view = this.props.view;
        let content;

        switch (view) {
            case "initial":
                content = (
                    <Blockquote>
                        Can we predict a GitHub repository's number of stars?
                    </Blockquote>
                );
                break;

            case 1:
                content = <div />;
                break;

            case "flowchart":
                content = <div>Another test</div>;
                break;

            default:
                content = null;
        }

        return <Flex>{content}</Flex>;
    }
}

export default SidePanel;
