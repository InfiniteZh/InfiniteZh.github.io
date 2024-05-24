import requests

# 定义 API 地址和身份验证 token
api_url = "http://47.99.99.38:8000/v1/chat/completions"
auth_token = "uVDIFaKG3FlLjvufSdeIpczLsb1cgwkJkof_BNpwU_TOTNChZBoeM1KJexdfb9zhYnsN5Zos6qISCrRt7mGxbigG2Cd4fWaCmBZHIzsgdZq64XXWQgyKFeuf0vpmV*s*CT58JlM_1t$w37U$bx8LPiGZ0"

# 设置请求头部
headers = {
    "Authorization": f"Bearer {auth_token}",
    "Content-Type": "application/json",
}

# 定义请求数据
data = {
    "model": "qwen",
    "messages": [
        {
            "role": "user",
            "content": "1+1=？"
        }
    ],
    "stream": False
}

# 发送 POST 请求
response = requests.post(api_url, json=data, headers=headers)

# 解析响应数据
if response.status_code == 200:
    response_data = response.json()
    completion_text = response_data["choices"][0]["message"]["content"]
    print("GPT 完成的对话内容:", completion_text)
else:
    print("请求失败:", response.text)
