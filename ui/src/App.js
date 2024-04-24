import React, { useState } from 'react';
import './App.css';

function fetchHeadline() {
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

  const {status, statusColor, subLine} = fetchHeadline();

  const [inConflict, setInConflict] = useState(false); // Default state is 'NO'
  const articles = [
    { title: "Article Title 1", description: "Detailed analysis of recent global events." },
    { title: "Article Title 2", description: "Expert opinions on geopolitical tensions." },
    { title: "Article Title 3", description: "Historical context behind today's conflicts." },
    { title: "Article Title 4", description: "Predictive insights into future global trends." }
  ];

  return (
    <div className="container">
      <div className="status-column">
        <h1 style={{color: statusColor}}>{status}</h1>
        <p>{subLine}</p>
        <button onClick={() => setInConflict(!inConflict)}>
          Toggle Conflict Status
        </button>
      </div>
      <div className="news-column">
        {articles.slice(0, 2).map((article, index) => (
          <div key={index} className="article">
            <h2>{article.title}</h2>
            <p>{article.description}</p>
          </div>
        ))}
      </div>
      <div className="news-column">
        {articles.slice(2).map((article, index) => (
          <div key={index} className="article">
            <h2>{article.title}</h2>
            <p>{article.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
