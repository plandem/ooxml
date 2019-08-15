package chart_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml/chart"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestChart(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
	<c:chart>
		<c:title>
			<c:overlay val="0"/>
			<c:spPr>
				<a:noFill/>
				<a:ln>
					<a:noFill/>
				</a:ln>
				<a:effectLst/>
			</c:spPr>
			<c:txPr>
				<a:bodyPr rot="0" spcFirstLastPara="1" vertOverflow="ellipsis" vert="horz" wrap="square" anchor="ctr" anchorCtr="1"/>
				<a:lstStyle/>
				<a:p>
					<a:pPr>
						<a:defRPr sz="1400" b="0" i="0" u="none" strike="noStrike" kern="1200" spc="0" baseline="0">
							<a:solidFill>
								<a:schemeClr val="tx1">
									<a:lumMod val="65000"/>
									<a:lumOff val="35000"/>
								</a:schemeClr>
							</a:solidFill>
							<a:latin typeface="+mn-lt"/>
							<a:ea typeface="+mn-ea"/>
							<a:cs typeface="+mn-cs"/>
						</a:defRPr>
					</a:pPr>
					<a:endParaRPr lang="en-US"/>
				</a:p>
			</c:txPr>
		</c:title>
		<c:autoTitleDeleted val="0"/>
		<c:plotArea>
			<c:layout/>
			<c:barChart>
				<c:barDir val="col"/>
				<c:grouping val="clustered"/>
				<c:varyColors val="0"/>
				<c:ser>
					<c:idx val="0"/>
					<c:order val="0"/>
					<c:spPr>
						<a:solidFill>
							<a:schemeClr val="accent1"/>
						</a:solidFill>
						<a:ln>
							<a:noFill/>
						</a:ln>
						<a:effectLst/>
					</c:spPr>
					<c:invertIfNegative val="0"/>
					<c:val>
						<c:numRef>
							<c:f>Sheet1!$A$1:$A$5</c:f>
							<c:numCache>
								<c:formatCode>General</c:formatCode>
								<c:ptCount val="5"/>
								<c:pt idx="0">
									<c:v>1</c:v>
								</c:pt>
								<c:pt idx="1">
									<c:v>2</c:v>
								</c:pt>
								<c:pt idx="2">
									<c:v>3</c:v>
								</c:pt>
								<c:pt idx="3">
									<c:v>4</c:v>
								</c:pt>
								<c:pt idx="4">
									<c:v>5</c:v>
								</c:pt>
							</c:numCache>
						</c:numRef>
					</c:val>
				</c:ser>
				<c:ser>
					<c:idx val="1"/>
					<c:order val="1"/>
					<c:spPr>
						<a:solidFill>
							<a:schemeClr val="accent2"/>
						</a:solidFill>
						<a:ln>
							<a:noFill/>
						</a:ln>
						<a:effectLst/>
					</c:spPr>
					<c:invertIfNegative val="0"/>
					<c:val>
						<c:numRef>
							<c:f>Sheet1!$B$1:$B$5</c:f>
							<c:numCache>
								<c:formatCode>General</c:formatCode>
								<c:ptCount val="5"/>
								<c:pt idx="0">
									<c:v>2</c:v>
								</c:pt>
								<c:pt idx="1">
									<c:v>4</c:v>
								</c:pt>
								<c:pt idx="2">
									<c:v>6</c:v>
								</c:pt>
								<c:pt idx="3">
									<c:v>8</c:v>
								</c:pt>
								<c:pt idx="4">
									<c:v>10</c:v>
								</c:pt>
							</c:numCache>
						</c:numRef>
					</c:val>
				</c:ser>
				<c:ser>
					<c:idx val="2"/>
					<c:order val="2"/>
					<c:spPr>
						<a:solidFill>
							<a:schemeClr val="accent3"/>
						</a:solidFill>
						<a:ln>
							<a:noFill/>
						</a:ln>
						<a:effectLst/>
					</c:spPr>
					<c:invertIfNegative val="0"/>
					<c:val>
						<c:numRef>
							<c:f>Sheet1!$C$1:$C$5</c:f>
							<c:numCache>
								<c:formatCode>General</c:formatCode>
								<c:ptCount val="5"/>
								<c:pt idx="0">
									<c:v>3</c:v>
								</c:pt>
								<c:pt idx="1">
									<c:v>6</c:v>
								</c:pt>
								<c:pt idx="2">
									<c:v>9</c:v>
								</c:pt>
								<c:pt idx="3">
									<c:v>12</c:v>
								</c:pt>
								<c:pt idx="4">
									<c:v>15</c:v>
								</c:pt>
							</c:numCache>
						</c:numRef>
					</c:val>
                </c:ser>
				<c:dLbls>
					<c:showLegendKey val="0"/>
					<c:showVal val="0"/>
					<c:showCatName val="0"/>
					<c:showSerName val="0"/>
					<c:showPercent val="0"/>
					<c:showBubbleSize val="0"/>
				</c:dLbls>
				<c:gapWidth val="219"/>
				<c:overlap val="-27"/>
				<c:axId val="753460511"/>
				<c:axId val="753462191"/>
			</c:barChart>
			<c:catAx>
				<c:axId val="753460511"/>
				<c:scaling>
					<c:orientation val="minMax"/>
				</c:scaling>
				<c:delete val="0"/>
				<c:axPos val="b"/>
				<c:majorTickMark val="none"/>
				<c:minorTickMark val="none"/>
				<c:tickLblPos val="nextTo"/>
				<c:spPr>
					<a:noFill/>
					<a:ln w="9525" cap="flat" cmpd="sng" algn="ctr">
						<a:solidFill>
							<a:schemeClr val="tx1">
								<a:lumMod val="15000"/>
								<a:lumOff val="85000"/>
							</a:schemeClr>
						</a:solidFill>
						<a:round/>
					</a:ln>
					<a:effectLst/>
				</c:spPr>
				<c:txPr>
					<a:bodyPr rot="-60000000" spcFirstLastPara="1" vertOverflow="ellipsis" vert="horz" wrap="square" anchor="ctr" anchorCtr="1"/>
					<a:lstStyle/>
					<a:p>
						<a:pPr>
							<a:defRPr sz="900" b="0" i="0" u="none" strike="noStrike" kern="1200" baseline="0">
								<a:solidFill>
									<a:schemeClr val="tx1">
										<a:lumMod val="65000"/>
										<a:lumOff val="35000"/>
									</a:schemeClr>
								</a:solidFill>
								<a:latin typeface="+mn-lt"/>
								<a:ea typeface="+mn-ea"/>
								<a:cs typeface="+mn-cs"/>
							</a:defRPr>
						</a:pPr>
						<a:endParaRPr lang="en-US"/>
					</a:p>
				</c:txPr>
				<c:crossAx val="753462191"/>
				<c:crosses val="autoZero"/>
				<c:auto val="1"/>
				<c:lblAlgn val="ctr"/>
				<c:lblOffset val="100"/>
				<c:noMultiLvlLbl val="0"/>
			</c:catAx>
			<c:valAx>
				<c:axId val="753462191"/>
				<c:scaling>
					<c:orientation val="minMax"/>
				</c:scaling>
				<c:delete val="0"/>
				<c:axPos val="l"/>
				<c:majorGridlines>
					<c:spPr>
						<a:ln w="9525" cap="flat" cmpd="sng" algn="ctr">
							<a:solidFill>
								<a:schemeClr val="tx1">
									<a:lumMod val="15000"/>
									<a:lumOff val="85000"/>
								</a:schemeClr>
							</a:solidFill>
							<a:round/>
						</a:ln>
						<a:effectLst/>
					</c:spPr>
				</c:majorGridlines>
				<c:numFmt formatCode="General" sourceLinked="1"/>
				<c:majorTickMark val="none"/>
				<c:minorTickMark val="none"/>
				<c:tickLblPos val="nextTo"/>
				<c:spPr>
					<a:noFill/>
					<a:ln>
						<a:noFill/>
					</a:ln>
					<a:effectLst/>
				</c:spPr>
				<c:txPr>
					<a:bodyPr rot="-60000000" spcFirstLastPara="1" vertOverflow="ellipsis" vert="horz" wrap="square" anchor="ctr" anchorCtr="1"/>
					<a:lstStyle/>
					<a:p>
						<a:pPr>
							<a:defRPr sz="900" b="0" i="0" u="none" strike="noStrike" kern="1200" baseline="0">
								<a:solidFill>
									<a:schemeClr val="tx1">
										<a:lumMod val="65000"/>
										<a:lumOff val="35000"/>
									</a:schemeClr>
								</a:solidFill>
								<a:latin typeface="+mn-lt"/>
								<a:ea typeface="+mn-ea"/>
								<a:cs typeface="+mn-cs"/>
							</a:defRPr>
						</a:pPr>
						<a:endParaRPr lang="en-US"/>
					</a:p>
				</c:txPr>
				<c:crossAx val="753460511"/>
				<c:crosses val="autoZero"/>
				<c:crossBetween val="between"/>
			</c:valAx>
			<c:spPr>
				<a:noFill/>
				<a:ln>
					<a:noFill/>
				</a:ln>
				<a:effectLst/>
			</c:spPr>
		</c:plotArea>
		<c:plotVisOnly val="1"/>
		<c:dispBlanksAs val="gap"/>
		<c:showDLblsOverMax val="0"/>
    </c:chart>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	drw := &chart.Chart{}
	err := decoder.DecodeElement(drw, nil)
	require.Nil(t, err)

	//encode data should be same as original
	encode, err := xml.Marshal(drw)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("c:", "").Replace(data), string(encode))
}

