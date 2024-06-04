data => {
    var xpath = "//text()";
    var result = document.evaluate(xpath, document, null, XPathResult.ANY_TYPE, null);
    var texts = [];
    var node = result.iterateNext();
    let title = document.title.trim();
    let link = document.location.href;
    let ifjump = false;

    let isEnglishMajority = str => {
        let englishCount = 0;
        let chineseCount = 0;
    
        // 遍历字符串
        for (let i = 0; i < str.length; i++) {
        const charCode = str.charCodeAt(i);
    
        // 判断字符是否为英文字符（基于ASCII码）
        if (charCode >= 65 && charCode <= 90 || charCode >= 97 && charCode <= 122) {
            englishCount++;
        }
        // 判断字符是否为中文字符（基于Unicode范围）
        else if (charCode >= 0x4e00 && charCode <= 0x9fa5) {
            chineseCount++;
        }
        }
    
        // 判断英文是否为大多数
        return englishCount > chineseCount;
    }
    while (node) {
        var parent = node.parentNode;
        if (parent && node.tagName != "style" && node.tagName != "script" && parent.tagName.toLowerCase() !== 'script' && parent.tagName.toLowerCase() !== 'style' && parent.tagName.toLowerCase() != "iframe") {
            if (node.textContent.trim().length > 5){
                let txt = node.textContent.trim();
                console.log("Dealing:",txt);
                texts.push(node.textContent.trim());
                ifjump = false;
            }else{
                if (! ifjump){
                texts.push("\n");
                ifjump = true;
                }
            }
        }
        node = result.iterateNext();
    }
    var allText = texts.join('\n');
    paragraphs = allText.split("\n\n")
    upload_txt = ""
    for(let line of paragraphs){
        if (upload_txt.length != 0){
            upload_txt += "\n"
        }
        let para = line.replace(/<[^>]*>.*<\/.*>/g, "")
        if (para.indexOf("<style ") > -1 ){
            continue
        }
        
        if (para.length < 20 && para.indexOf(" ")< 0){
            continue
        }
        if (para.trim().indexOf("\n") > -1){
            for(let line of para.split("\n")){
                if (isEnglishMajority(line)){
                    let words = line.trim().split(" ")
                    if (words.length < 3){
                        continue
                    }else{
                        upload_txt += "\n"+line.trim()
                    }
                }else{
                    if (line.length < 20 ){
                        continue
                    }else{
                        upload_txt += "\n"+line.trim()
                    }
                }
                
            }
        }else{
            
            if (para.length > 5){
                upload_txt += "\n"+para
            }        
        }
    }
    return upload_txt
}
