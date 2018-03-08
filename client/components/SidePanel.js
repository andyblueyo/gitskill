import React from "react";

import Flex from "./Flex";
import Blockquote from "./Blockquote";
import GitHubStar from "./GitHubStar";
import Terminal from "./Terminal";

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

            case "terminal":
                content = <Terminal />;
                break;

            default:
                content = null;
        }

        return <Flex>{content}</Flex>;
    }
}

export default SidePanel;
