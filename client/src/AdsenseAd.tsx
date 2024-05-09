import { useEffect } from 'react';

export default function AdsenseAd() {
    useEffect(() => {
        try {

            console.log("Ad placed!");
            ((window as any).adsbygoogle = (window as any).adsbygoogle || []).push({});

        } catch (err) {
            console.error(err);
        }
    }, []);

    return (
        <ins className="adsbygoogle m-1"
            style={{ display: "block" }}
            data-ad-client="ca-pub-5382375998819941"
            data-ad-slot="4503418768"
            data-ad-format="auto"
            data-full-width-responsive="true"></ins>
    );
};

