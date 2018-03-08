import React from "react";

import Flex from "./Flex";
import Blockquote from "./Blockquote";
import GitHubStar from "./GitHubStar";

class SidePanel extends React.Component {
    render() {
        let view = this.props.view;
        let content;

        switch (view) {
            case "initial":
                content = (
                    <div>
                        <Blockquote>
                            Can we predict a GitHub repository&#39;s number of
                            stars?
                        </Blockquote>
                        <GitHubStar />
                    </div>
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
