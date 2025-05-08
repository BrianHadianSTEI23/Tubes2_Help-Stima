import requests
from bs4 import BeautifulSoup
import csv

url = "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"
response = requests.get(url)
soup = BeautifulSoup(response.content, "html.parser")

tables = soup.find_all("table", class_="list-table col-list icon-hover")

elements_data = []
elements_data.append(["Air", "-", "-"])
elements_data.append(["Earth", "-", "-"])
elements_data.append(["Water", "-", "-"])
elements_data.append(["Fire", "-", "-"])
elements_data.append(["Time", "-", "-"])

for table in tables:
    rows = table.find_all("tr")[1:]
    last_element_name = ""

    for row in rows:
        cols = row.find_all("td")

        if len(cols) == 2:
            element_name = cols[0].text.strip()
            last_element_name = element_name
            combo_cell = cols[1]
        else:
            continue

        list_items = combo_cell.find_all("li")
        if list_items:  # kalau ada beberapa
            combinations = [li.text.strip() for li in list_items]
        else:
            raw_text = combo_cell.get_text(strip=True)
            if raw_text:
                combinations = [raw_text]
            else:
                continue

        for combination in combinations:
            if '+' in combination:
                parts = combination.split('+')
                ingredient1 = parts[0].strip()
                ingredient2 = parts[1].strip()
            else:
                ingredient1 = combination.strip()
                ingredient2 = ""

            if(len(ingredient1) <= 20): elements_data.append([element_name, ingredient1, ingredient2])

# Simpan ke CSV
filename = "./little_alchemy_2_elements_split.csv"
with open(filename, mode="w", newline="", encoding="utf-8") as file:
    writer = csv.writer(file, delimiter=";")
    writer.writerow(["Element", "Ingredient1", "Ingredient2"])
    writer.writerows(elements_data)

print(f"Sukses! Data disimpan di '{filename}' dengan {len(elements_data)} kombinasi.")
