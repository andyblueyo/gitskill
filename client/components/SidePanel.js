import React from "react";

import Flex from "./Flex";
import Blockquote from "./Blockquote";
import GitHubStar from "./GitHubStar";
import Terminal from "./Terminal";
import TableComponent from "./TableComponent";

class SidePanel extends React.Component {
    render() {
        let view = this.props.view;
        let content;

        switch (view) {
            case "initial":
                content = (
                    <div>
                        <Blockquote>
                            Can we predict a GitHub repository's number of
                            stars?
                        </Blockquote>
                        <GitHubStar />
                    </div>
                );
                break;

            case "terminal":
                content = <Terminal />;
                break;

            case "prePrepData":
                content = (
                    <TableComponent
                        columns={[
                            "userType",
                            "publicRepos",
                            "orgs",
                            "ownerType",
                            "stars",
                            "forks",
                            "Java"
                        ]}
                        data={this.props.prePrepData}
                    />
                );
                break;

            default:
                content = null;
        }

        return <Flex>{content}</Flex>;
    }
}

export default SidePanel;
