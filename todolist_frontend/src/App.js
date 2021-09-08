import React, { useState } from 'react'
import './App.css';
import Form from './components/Form';

function App() {
  const [inputText, setInputText] = useState("");

  return (
    <div>
      <Form inputText={inputText} setInputText={setInputText} />
    </div>
    
  );
}

export default App;
