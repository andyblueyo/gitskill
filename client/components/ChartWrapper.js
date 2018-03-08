//     [var name:"feature_imporantces" value:"./data/feature_importances.json" /]
//    [var name:"clean_f_i" value:`feature_imporantces.map((f) => {x:f.feature, y:f.importance})` /]
import React from "react";
import Chart from "./default/chart.js";
import d from "../data/feature_importances.json"
import {VictoryTooltip} from "victory";

const colorScale = [
  "#0066FF",
  "#3385FF",
  "#66A3FF",
  "#99C2FF",
  "#CCE0FF",
  "#FFFFFF",
]

class ChartWrapper extends React.Component {
    componentDidMount() {
        // console.log(this.props.data);
    }
    render() {
        let data = d.filter((f, i) => i < 4).map((f) => {
          return {x:f.feature, y:f.importance, label:f.feature}
        });
        return (
            <div style={{width: "300px"}}>
                <Chart
                 labelComponent={<VictoryTooltip/>}
                 colorScale={colorScale}
                 data={data}
                 type={this.props.type}
                 events={[{
                  target: "data",
                  eventHandlers: {
                    onMouseOver: () => {
                      return [
                        {
                          target: "data",
                          mutation: () => ({style: {fill: "gold", width: 30}})
                        }, {
                          target: "labels",
                          mutation: () => ({ active: true })
                        }
                      ];
                    },
                    onMouseOut: () => {
                      return [
                        {
                          target: "data",
                          mutation: () => {}
                        }, {
                          target: "labels",
                          mutation: () => ({ active: false })
                        }
                      ];
                    }
                  }
                }]}
               />
            </div>
        );
    }
}

export default ChartWrapper;
