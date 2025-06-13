import React from "react";
import { env } from "@/env";
import { QRCodeSVG } from "qrcode.react";

const STRAVA_LOGIN_URL = `http://www.strava.com/oauth/authorize?client_id=${env.VITE_STRAVA_CLIENT_ID}&response_type=code&redirect_uri=${env.VITE_SERVER_URL}/strava/exchange-token&scope=profile:read_all,activity:read_all`;

const Login: React.FC = () => {
  return (
    <div className="space-y-4">
      <p className="text-xl">Please log in to see your Strava stats.</p>

      <a
        href={STRAVA_LOGIN_URL}
        target="_blank"
        rel="noopener noreferrer"
        className="mb-32 block"
      >
        {env.VITE_IS_PROD === false ? (
          <>Login</>
        ) : (
          <QRCodeSVG value={STRAVA_LOGIN_URL} size={280} className="inline" />
        )}
      </a>
    </div>
  );
};

export default Login;
