import AdsenseAd from "./AdsenseAd"
import Content from "./Content"

export default function App() {

  return (
    <>
      <div className="flex flex-row items-center">
        {/* LEFT SIDE ADS */}
        <div className="w-[400px] px-[50px]">
          <AdsenseAd />
          <AdsenseAd />
        </div>
        {/* CONTENT */}
        <div className="flex-1">
          <Content />
        </div>
        {/* RIGHT SIDE ADS */}
        <div className="w-[400px] px-[50px]">
          <AdsenseAd />
          <AdsenseAd />
        </div>
      </div>
    </>
  )
}