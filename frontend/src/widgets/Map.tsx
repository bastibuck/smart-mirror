import React from "react";

import MapLibre, { Layer, Source } from "react-map-gl/maplibre";
import "maplibre-gl/dist/maplibre-gl.css";
import polyline from "@mapbox/polyline";

import WidgetPositioner from "./_layout/WidgetPositioner";

// TODO: get dedicated endpoint for /last-activity
// TODO: parse polyline in server
// TODO: show more info about last activity over map

const Map: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const decryptedCoordinates = polyline
    .decode(
      "ozrjIio{|@\\G`AEN@HD`@BlACd@D\\RTVd@Tb@`@RL`@TZJZTf@Tl@`@l@ZZTVHLJXDh@^VBN?h@Ud@IJ@DBFVLtAX|@DhADRHHNIh@g@TOZINI\\A`@Gn@OZCHHB`@Pv@HbAPf@HfAJl@Fp@LXDDV?t@Hf@HbAg@TQBEDS?g@[qCg@uCAQIk@M_CGkCMeAI[MUg@k@YMQAc@?YIIGg@}AQ_AMSO}@@o@PyB?oAF}AE{@@a@BIJGh@NX@\\HPANEFK@QASMa@MQSUi@]}@u@UK{@{@c@HE?OMG[E]AcBS}Aa@oAc@u@yAy@i@]Yi@]e@{@_AYUc@m@YUQYYWQ_@OKSWg@i@e@o@YSs@u@SW]W_Ay@e@[[i@e@ScAy@g@q@SOa@K]_@g@Ui@_@_@MQM[c@Q]MMQWYOaB_B_@u@e@s@Yy@m@_AM[o@eAmAaCy@eA[UQ]q@u@W_@}@}@a@WUGe@HOAeAHa@HIDa@Cs@FKAm@Rs@BuATKIM[KGE?q@Hu@NO?e@FaAN]Ci@Rm@P[PK@u@d@]Fa@XQT]Nc@\\c@TwAz@_DlAOJUFMLc@pASRGLYVSV}@pAWVoCnDw@|@OTMZw@v@{@pA]z@W\\_@n@c@dA_@nAUb@a@dAa@tAo@vAy@fCi@pAOh@[t@Kf@c@nAGLHd@JVx@lAf@|@t@~Av@t@JZZb@XZT`@Zb@FLHj@DLJHBAb@K\\Af@DPHNLLP^rAJPNNXNNJLPT`@HDHBLCz@o@l@[vBi@p@YJABCf@KTAd@StB_@nAG|@N|@@tBK\\?^Mt@JjAAp@C^BdB?z@Dz@Nv@Rx@V^R^JJJfA`CFf@Nt@?\\Fl@GdEBPHPFDj@Hd@?~@K",
    )
    // flip order
    .map(([lat, lng]) => [lng, lat]);

  const bounds = getMaxBounds(decryptedCoordinates);

  return (
    <>
      <WidgetPositioner {...widgetPositionerProps}>
        <MapLibre
          onLoad={(e) => {
            e.target.fitBounds(bounds, {
              animate: false,
            });
          }}
          interactive={false}
          attributionControl={{ compact: false }}
        >
          <Source
            type="geojson"
            data={{
              type: "Feature",
              properties: null,
              geometry: {
                type: "LineString",
                coordinates: decryptedCoordinates,
              },
            }}
          >
            <Layer type="line" paint={{ "line-color": "white" }} />
          </Source>
        </MapLibre>
      </WidgetPositioner>
    </>
  );
};

export default Map;

/**
 * calculates maximum bounds from a list of [lng, lat]
 *
 * @returns a list of max bounds with order [west, south, east, north]
 */
function getMaxBounds(coords: number[][]) {
  const lats = coords.map(([, lat]) => lat);
  const lngs = coords.map(([lng]) => lng);

  const minLat = Math.min(...lats);
  const maxLat = Math.max(...lats);
  const minLng = Math.min(...lngs);
  const maxLng = Math.max(...lngs);

  return [minLng, minLat, maxLng, maxLat] as [number, number, number, number];
}
