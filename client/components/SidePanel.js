import React from "react";

import Flex from "./Flex";
import Blockquote from "./Blockquote";
import GitHubStar from "./GitHubStar";
import Terminal from "./Terminal";
import TableComponent from "./TableComponent";
import Chart from "./Chart";

class SidePanel extends React.Component {
    render() {
        let view = this.props.view;
        let content;

        switch (view) {
            case "initial":
                content = (
                    <div>
                        <Blockquote>
                            Based on data from GitHub, are we able to predict
                            the number of stars a repo from an organization or
                            user can have?
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

            case "predAct":
                content = (
                    <Chart
                        src="images/predAct.svg"
                        figureNumber={3}
                        figureDescription="Integer posuere erat a ante venenatis dapibus posuere velit aliquet"
                    />
                );
                break;

            case "predError":
                content = (
                    <Chart
                        src="images/predError.svg"
                        figureNumber={4}
                        figureDescription="Morbi leo risus, porta ac consectetur ac, vestibulum at eros"
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
