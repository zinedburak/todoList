import React, { useState } from "react"
import {render,fireEvent, queryByPlaceholderText} from '@testing-library/react'
import Form from "../Form"
import "@testing-library/jest-dom/extend-expect"


it("renders correctly", () => {
    
    const component = render(<Form/>)
    const button = component.getByTestId("add-button");
    const input_text = component.getByPlaceholderText("Add Todo");
    expect(button).toBeTruthy;
    expect(input_text).toBeTruthy;
    

});


describe("Input Value", () => {
    it("updates on change", () => {
        const setInputText = jest.fn();
        const component = render(<Form setInputText={setInputText}/>)
        const todoInput = component.getByPlaceholderText("Add Todo")

        fireEvent.change(todoInput, {target: {value: "test"}})

        expect(todoInput.value).toBe("test")
    })
})

describe("Add Button" , () => {
    describe("calls function when pressed with no input text" , () => {
        it("it does not trigger anything" , () => {
            const setInputText = jest.fn();
            const component = render(<Form setInputText={setInputText}/>)
            const todoInput = component.getByPlaceholderText("Add Todo")
            const addButton = component.getByTestId("add-button")

            fireEvent.click(addButton)

            expect(setInputText).not.toHaveBeenCalled()
        })
    })
    describe("calls function when pressed with input data" ,() => {
        it("it does trigger anything" , () => {
            const setInputText = jest.fn();
            const component = render(<Form setInputText={setInputText}/>)
            const todoInput = component.getByPlaceholderText("Add Todo")
            const addButton = component.getByTestId("add-button")

            fireEvent.change(todoInput, {target: {value: "test"}})
            fireEvent.click(addButton)

            expect(setInputText).toHaveBeenCalled()
        })
    })
})