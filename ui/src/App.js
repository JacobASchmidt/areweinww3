import React, { useEffect, useState } from 'react';
import './App.css';

function fetchStatus() {
  return fetch("api/v1/status").then((resp) => resp.json());
  const statuses = [
    {
      status: "YES",
      statusColor: "#e63946",
      subLine: "Global Conflict",
      explanation: "Quite a long drawn out explanation about how israel nuked iran or something"
    },
    {
      status: "NO",
      statusColor: "#4caf50",
      subLine: "Regional Wars",
      explanation: "Quite a long drawn out explanation about how israel hasn't nuked iran or something"
    },
  ];
  return statuses[Math.floor(Math.random() * statuses.length)]
}

function App() {

  const [status, setStatus] = useState({status: '', textColor: '', subLine: '', explanation: ''});

  useEffect(() => {
    fetchStatus().then(setStatus);
  }, []);


  const articles = [
    { headline: "Article Title 1", description: "Detailed analysis of recent global events." },
    { headline: "Article Title 2", description: "Expert opinions on geopolitical tensions." },
    { headline: "Article Title 3", description: "Historical context behind today's conflicts." },
    { headline: "Article Title 4", description: "Predictive insights into future global trends." }
  ];

  return (
    <div className="container">
      <div className="status-column">
        <h1 style={{color: status.textColor}}>{status.status}</h1>
        <h2>{status.subLine}</h2>
        <p>
          {status.explanation}
        </p>
      </div>
      <div className="news-column">
        {articles.slice(0, 2).map((article, index) => (
          <div key={index} className="article">
            <h2>{article.headline}</h2>
            <p>{article.description}</p>
          </div>
        ))}
      </div>
      <div className="news-column">
        {articles.slice(2).map((article, index) => (
          <div key={index} className="article">
            <h2>{article.headline}</h2>
            <p>{article.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
