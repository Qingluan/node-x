data => {
    let Xpath = (path, ele=null, filter=null)=>{
        if (ele == null){
            ele = document;
        }
        var result = document.evaluate(path, ele, null, XPathResult.ANY_TYPE, null);
        var nodes = [];
        var node = result.iterateNext();
        // 遍历所有匹配的节点
        while (node) {
            // 检查父节点是否是script或style标签
            var parent = node.parentNode;
            if (parent && parent.tagName.toLowerCase() !== 'script' && parent.tagName.toLowerCase() !== 'style') {
                // 如果不是，则将文本节点的内容添加到数组中
                if (node.textContent.trim().length > 0){
                    if (filter != null && typeof(filter) == "function"){
                        if (filter(node)){
                            nodes.push(node);
                        }
                    }else{
                        nodes.push(node);
                    }
                }else{
                }
            }
            node = result.iterateNext();
        }
        return nodes
    }

    let eles = Xpath("//div/span/a/h3")
    let searchItems = [];
    for(let h3 of eles){
        let title = h3.textContent.trim();
        let url = h3.parentElement.href;
        let desc = h3.parentElement.parentElement.textContent.trim().replace(url,"").replace(title,"").trim();
        let item = {
            title: title,
            url: url,
            desc: desc
        }
        searchItems.push(item);
    }
    return JSON.stringify(searchItems);
}