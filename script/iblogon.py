from selenium import webdriver
import os
import time 

startWebPortal = './bin/run.sh ./root/config.yml &'
os.system(startWebPortal)
time.sleep(4)

driver = webdriver.Chrome()

driver.get("https://localhost:5000")
driver.find_element_by_id("user_name").send_keys("phage717")
driver.find_element_by_id("password").send_keys("tin25GAP75")
driver.find_element_by_id("submitForm").click()

