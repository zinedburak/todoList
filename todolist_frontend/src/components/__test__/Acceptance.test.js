import puppeteer from 'puppeteer'
import "@testing-library/jest-dom/extend-expect"
import "@testing-library/jest-dom"



    it("Passes Acceptance Test", async () => {
        jest.setTimeout(10000);
        const browser = await puppeteer.launch();
        
        try{
        const page = await browser.newPage();
        await page.goto("http://localhost:3000/");
        
        await page.click('#input_text');
        await page.type('#input_text', 'test1');
        await page.click('#add-button') 
        await page.waitForResponse(res => res.url().endsWith('add_todos'))
        await page.waitForResponse(res => res.url().endsWith('get_todos'))
        const todoItem = await page.$('#todo-item')
        const todoItemValue = await todoItem.evaluate((el) => el.textContent)

        expect(todoItemValue).toContain("test1")
        }
        finally{
            browser.close();
        }
    })



