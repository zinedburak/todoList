import React, { useState } from "react"

const Form = ({inputText, setInputText}) =>{
    const inputTextHandler = (e) => {
        setInputText(e.target.value);
    }
    const submitTodoHandler = (e) => {
        e.preventDefault();
        if(inputText){
            setInputText("1")
        }
    }

    return (
        <form>
            <input value={inputText} onChange={inputTextHandler} type="text" className="todo-input" placeholder="Add Todo" />
            <button onClick={submitTodoHandler} className="todo-button" type="submit" data-testid="add-button">
                <i className="fas fa-plus-square"></i>
            </button>
        </form>
    )
}

export default Form;