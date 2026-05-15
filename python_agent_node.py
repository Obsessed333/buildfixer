import os
import requests
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from openai import OpenAI

app = FastAPI(title="PoB AI Fixer Agent Node")
client = OpenAI(
    base_url="https://openrouter.ai",
    api_key=os.getenv("OPENROUTER_API_KEY")
)

class AnalyzeRequest(BaseModel):
    xml_data: str
    budget: str

@app.post("/analyze")
async def analyze_character(payload: AnalyzeRequest):
    try:
        
        extract_response = requests.post(
            "http://127.0.0.1:8001/extract", 
            json={"xml_data": payload.xml_data}
        )
        if extract_response.status_code != 200:
            raise HTTPException(status_code=400, detail="Data extraction processing failed.")
            
        parsed_data = extract_response.json()
        
      
        metrics = parsed_data["metrics"]
        gear = parsed_data["equipped_gear"]
        
        
        system_instructions = (
            "You are an expert Path of Exile build engineer. Review the character statistics "
            "and gear. Locate under-performing items and recommend explicit modifiers to purchase."
        )
        
        user_prompt = f"""
        Class: {metrics['class']} (Level {metrics['level']})
        Allocation Budget: {payload.budget}
        
        Computed Character Stats:
        {metrics}
        
        Current Equipped Gear:
        {gear}
        """
        
       
        ai_response = client.chat.completions.create(
            model="meta-llama/llama-3-8b-instruct:free",
            messages=[
                {"role": "system", "content": system_instructions},
                {"role": "user", "content": user_prompt}
            ]
        )
        if isinstance(ai_response, str):
            analysis_text = ai_response
        elif hasattr(ai_response, 'choices') and len(ai_response.choices) > 0:
            analysis_text = ai_response.choices[0].message.content
        else:
            
            try:
                analysis_text = ai_response['choices'][0]['message']['content']
            except:
                analysis_text = str(ai_response)

        return {
            "status": "success",
            "stats": metrics,
            "analysis": analysis_text
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000) 
