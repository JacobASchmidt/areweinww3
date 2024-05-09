import { useEffect, useState } from "react"

interface Status {
  status: string;
  subline: string;
  explanation: string;
}
interface Article {
  author: string;
  title: string;
  description: string;
  url: string;
  urlToImage: string;
  content: string;
}

export default function Content() {
  const [status, setStatus] = useState<Status>();
  const [articles, setArticles] = useState<Article[]>();

  const fetchStatus = async () => {
    const response = await fetch(
      "/api/v1/status"
    );
    if (!response.ok) {
      console.error("Failed to fetch status");
      console.log(await response.text());
      return;
    }
    const data = await response.json();
    console.log("Status: ", data.status);
    console.log("Articles: ", data.articles);
    setStatus(data.status);
    setArticles(data.articles);
  };

  useEffect(() => {
    fetchStatus();
  }, []);

  return (
    <>
      <div className='flex flex-col items-center'>
        <div className="text-[50px]">Are we in World War 3?</div>
        <div className={"text-[100px] font-bold " + (status?.status == "YES" ? "text-red-500" : "text-green-500")} >
          {status?.status == "YES" ? "YES" : "NO"}
        </div>
        {/* ARTICLES */}
        {articles?.map((article, index) => (
          <div key={index} className="card card-side border border-base-content w-[800px] mb-[10px]" onClick={() => {
            window.open(article.url, '_blank');
          }}>
            <figure><img src={article.urlToImage} alt="Movie" /></figure>
            <div className="card-body">
              <h2 className="card-title">{article.title}</h2>
              <p>{article.content}</p>
            </div>
          </div>
        ))}


      </div>
    </>
  )
}