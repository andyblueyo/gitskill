import React from "react";

import Flex from "./Flex";
import Blockquote from "./Blockquote";
import GitHubStar from "./GitHubStar";
import Terminal from "./Terminal";
import TableComponent from "./TableComponent";
import Chart from "./Chart";
import Image from "./Image";

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
                        figureDescription="Our team member, Evan, wrote a script that would allow us to call data from the API using a Round Robin API token scraping method"
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

            case "langFrequency":
                content = (
                    <Image
                        src="images/langFrequency.png"
                        figureNumber={42}
                        figureDescription="Test"
                    />
                );
                break;

            case "predAct":
                content = (
                    <Image
                        src="images/predAct.png"
                        figureNumber={42}
                        figureDescription="Test"
                    />
                );
                break;

            case "predError":
                content = (
                    <Image
                        src="images/predError.png"
                        figureNumber={42}
                        figureDescription="Test"
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
