import React from "react";
import SVG from "./default/svg";

import Caption from "./Caption";
import Text from "./Text";

class Chart extends React.Component {
    render() {
        return (
            <div className="svg-wrapper">
                <SVG
                    src={this.props.src}
                    preloader={<Text>Loading chart...</Text>}
                />
                <Caption
                    figureNumber={this.props.figureNumber}
                    figureDescription={this.props.figureDescription}
                />
            </div>
        );
    }
}

export default Chart;
