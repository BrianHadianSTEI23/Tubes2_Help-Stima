"use client";

import dynamic from "next/dynamic";
import { useEffect, useRef, useState } from "react";

// Dynamically import the Tree component from react-d3-tree
const Tree = dynamic(() => import("react-d3-tree").then((mod) => mod.Tree), {
  ssr: false,
});

export default function RecipeTree({ data }) {
  const treeRef = useRef(null);
  const [dimensions, setDimensions] = useState({ width: 600, height: 600 });
  const [elementsImage, setElementsImage] = useState({});

  // Fetch the elements image data from the JSON file
  useEffect(() => {
    const fetchImageData = async () => {
      const response = await fetch("/icon/elements.json");
      const data = await response.json();
      setElementsImage(data.Images);
    };

    fetchImageData();
  }, []);

  // Calculate tree dimensions on component mount
  useEffect(() => {
    if (treeRef.current) {
      const { width, height } = treeRef.current.getBoundingClientRect();
      setDimensions({ width, height });
    }
  }, []);

  // Unified color for all nodes
  const nodeColor = "#B0B0B0";  // Set to a pleasant color

  // Memoize the custom node renderer to prevent unnecessary re-renders
  const renderCustomNode = ({ nodeDatum }) => {
    const name = String(nodeDatum.name);
    const paddingX = 20;
    const charWidth = 8;
    const textWidth = name.length * charWidth;
    const boxWidth = textWidth + paddingX + 20;
    const boxHeight = 60; // Increased height for accommodating the image and text

    // Get the image link for the current node (element)
    const imageLink = elementsImage[name] || "";

    return (
      <g transform="rotate(180)">
        <rect
          x={-boxWidth / 2}
          y={-boxHeight / 2}
          width={boxWidth}
          height={boxHeight}
          fill={nodeColor}
          stroke="#191919"
          strokeWidth="1"
          rx={8}
        />
        
        {/* Display the image above the text and center it */}
        {imageLink && (
          <image
            href={imageLink}
            x={-15}  // Positioning the image in the center (boxWidth / 2 - imageWidth / 2)
            y={-boxHeight / 2 + 5}  // Position the image slightly above the node
            width={30}  // Set a fixed width for the image
            height={30} // Set a fixed height for the image
          />
        )}

        {/* Display the text below the image */}
        <text
          fill="#000000" // Set the text color to black for better contrast
          x={0}
          y={boxHeight / 2 - 10} // Move the text slightly lower to leave space for the image
          textAnchor="middle"
          stroke="none"
          style={{
            fontSize: "16px", // Slightly larger text for better readability
            fontFamily: "Roboto, sans-serif",
            fontWeight: "700",  // A bit of boldness for clarity
          }}
        >
          {name}
        </text>
      </g>
    );
  };

  return (
    <div
      ref={treeRef}
      style={{
        width: "100%",
        height: "600px",
        transform: "rotate(180deg)",
        background: "#F9FAFB",
        padding: "1.5rem",  // Increased padding for better spacing
        borderRadius: "8px",
        boxSizing: "border-box",
        overflow: "hidden",
      }}
    >
      <Tree
        data={data}
        orientation="vertical"
        translate={{ x: dimensions.width/2, y: dimensions.height/2 }}
        pathFunc="elbow"
        depthFactor={120}
        collapsible={false}
        nodeSize={{ x: 200, y: 100 }}
        renderCustomNodeElement={renderCustomNode}
      />
    </div>
  );
}
