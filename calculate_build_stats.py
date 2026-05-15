import base64
import zlib
import xml.etree.ElementTree as ET
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI(title="PoB Data Extraction Node")

class ExtractRequest(BaseModel):
    xml_data: str

@app.post("/extract")
async def extract_metrics(payload: ExtractRequest):
    try:
        #Transmision token boundaries
        token = payload.xml_data.strip()
        if token.startswith('"') and token.endswith('"'):
            token = token[1:-1]
            
        token = token.replace('-', '+').replace('_', '/')
        missing_padding = len(token) % 4
        if missing_padding:
            token += '=' * (4 - missing_padding)

        # decompress Path of Building base64 string directly back into true XML bytes
        compressed_bytes = base64.b64decode(token)
        xml_bytes = zlib.decompress(compressed_bytes)
        
        # native XML parsing
        root = ET.fromstring(xml_bytes)
        
        build_node = root.find("Build")
        if build_node is None:
            raise ValueError("Invalid PoB XML layout data structure.")
            
        build_metrics = {
            "class": build_node.attrib.get("className", "Unknown"),
            "level": int(build_node.attrib.get("level", "1")),
            "total_dps": 0.0,
            "life": 0,
            "ehp": 0,
            "fire_res": 0,
            "cold_res": 0,
            "lightning_res": 0,
            "chaos_res": 0
        }
        
        # pull calculated parameters straight out of structural PlayerStat tags
        for player_stat in build_node.findall("PlayerStat"):
            stat_name = player_stat.attrib.get("stat")
            try:
                stat_val = float(player_stat.attrib.get("value", "0"))
                if stat_name == "TotalDPS":
                    build_metrics["total_dps"] = stat_val
                elif stat_name == "Life":
                    build_metrics["life"] = int(stat_val)
                elif stat_name == "EffectiveHP":
                    build_metrics["ehp"] = int(stat_val)
                elif stat_name == "FireResist":
                    build_metrics["fire_res"] = int(stat_val)
                elif stat_name == "ColdResist":
                    build_metrics["cold_res"] = int(stat_val)
                elif stat_name == "LightningResist":
                    build_metrics["lightning_res"] = int(stat_val)
                elif stat_name == "ChaosResist":
                    build_metrics["chaos_res"] = int(stat_val)
            except ValueError:
                continue

        # extract item dictionaries mapped by item ID
        items_node = root.find("Items")
        item_id_map = {}
        if items_node is not None:
            for item in items_node.findall("Item"):
                item_id = int(item.attrib.get("id", "0"))
                item_id_map[item_id] = item.text.strip() if item.text else ""

        # locate equipped items within the active item configuration setup
        equipped_gear = {}
        if items_node is not None:
            active_set_id = int(items_node.attrib.get("activeItemSet", "1"))
            active_set = None
            
            for item_set in items_node.findall("ItemSet"):
                if int(item_set.attrib.get("id", "1")) == active_set_id:
                    active_set = item_set
                    break
            
            if active_set is None:
                item_sets = items_node.findall("ItemSet")
                if item_sets:
                    active_set = item_sets[0]

            if active_set is not None:
                for slot in active_set.findall("Slot"):
                    slot_name = slot.attrib.get("name")
                    item_id = int(slot.attrib.get("itemId", "0"))
                    if item_id in item_id_map and item_id_map[item_id]:
                        equipped_gear[slot_name] = item_id_map[item_id]

        return {
            "status": "success",
            "metrics": build_metrics,
            "equipped_gear": equipped_gear
        }
        
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Native XML calculation failure: {str(e)}")

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8001)