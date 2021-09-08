import React from "react"
import {render,fireEvent, queryByPlaceholderText} from '@testing-library/react'
import Form from "../Form"
import "@testing-library/jest-dom/extend-expect"


it("renders correctly", () => {

    const component = render(<Form />);
    const button = component.getByTestId("add-button");
    const input_text = component.getByPlaceholderText("Add Todo");
    expect(button).toBeTruthy;
    expect(input_text).toBeTruthy;
    

});


// describe("Input Value", () => {
//     it("updates on change", () => {
//         const {queryByPlaceHolderText} = render([Form])

//         const todoInput = queryByPlaceHolderText("Add Todo")

//         fireEvent.change(todoInput, {target: {value: "test"}})

//         expect(todoInput.value).toBe("test")
//     })
// })