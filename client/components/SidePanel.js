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
                content = (
                    <Terminal
                        figureNumber={1}
                        figureDescription="Cras mattis consectetur purus sit amet fermentum"
                    />
                );
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
                        figureNumber={2}
                        figureDescription="Ullamcorper nulla non metus auctor fringilla"
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
